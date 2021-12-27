const colors = require("tailwindcss/colors");
const defaultTheme = require("tailwindcss/defaultTheme");

module.exports = {
  content: ["./www/**/*.{html,js}"],
  theme: {
    colors,
    extend: {
      fontFamily: {
        serif: ["IBM Plex Sans", ...defaultTheme.fontFamily.serif],
        sans: ["Inter var", ...defaultTheme.fontFamily.sans],
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
}
