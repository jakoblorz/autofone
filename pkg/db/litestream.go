package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/benbjohnson/litestream"
	"github.com/benbjohnson/litestream/gcs"
	"github.com/bep/debounce"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Instance struct {
	*sqlx.DB
	ctx    context.Context
	cancel context.CancelFunc

	lsdb *litestream.DB
	debc func(func())
}

func (i *Instance) init() {
	i.ctx, i.cancel = context.WithCancel(context.Background())
	i.debc = debounce.New(100 * time.Millisecond)
}

func (i *Instance) replicate() {
	i.lsdb.Replicas[0].Sync(context.Background())
}

func (i *Instance) Close() (err error) {
	if i.DB != nil {
		err = i.DB.Close()
	}
	if i.lsdb != nil {
		i.lsdb.SoftClose()
	}
	return
}

func (i *Instance) Commit(ctx context.Context) (err error) {
	err = i.lsdb.Sync(ctx)
	if err != nil {
		return
	}
	i.debc(i.replicate)
	return
}

func (i *Instance) GCP(ctx context.Context, dsn string, bucket string) (*Instance, error) {
	var err error

	i.init()

	replica := gcs.NewReplicaClient()
	replica.Bucket = bucket

	i.lsdb, err = replicate(ctx, dsn, replica, "gcs")
	if err != nil {
		return nil, err
	}

	i.DB, err = sqlx.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func replicate(ctx context.Context, dsn string, client litestream.ReplicaClient, name string) (*litestream.DB, error) {
	// Create Litestream DB reference for managing replication.
	lsdb := litestream.NewDB(dsn)

	replica := litestream.NewReplica(lsdb, name)
	replica.Client = client

	lsdb.Replicas = append(lsdb.Replicas, replica)

	if err := restore(ctx, replica); err != nil {
		return nil, err
	}

	// Initialize database.
	if err := lsdb.Open(); err != nil {
		return nil, err
	}

	return lsdb, nil
}

func restore(ctx context.Context, replica *litestream.Replica) (err error) {
	// Skip restore if local database already exists.
	if _, err := os.Stat(replica.DB().Path()); err == nil {
		fmt.Println("local database already exists, skipping restore")
		return nil
	} else if !os.IsNotExist(err) {
		return err
	}

	// Configure restore to write out to DSN path.
	opt := litestream.NewRestoreOptions()
	opt.OutputPath = replica.DB().Path()
	opt.Logger = log.New(os.Stderr, "", log.LstdFlags|log.Lmicroseconds)

	// Determine the latest generation to restore from.
	if opt.Generation, _, err = replica.CalcRestoreTarget(ctx, opt); err != nil {
		return err
	}

	// Only restore if there is a generation available on the replica.
	// Otherwise we'll let the application create a new database.
	if opt.Generation == "" {
		fmt.Println("no generation found, creating new database")
		return nil
	}

	fmt.Printf("restoring replica for generation %s\n", opt.Generation)
	if err := replica.Restore(ctx, opt); err != nil {
		return err
	}
	fmt.Println("restore complete")
	return nil
}
