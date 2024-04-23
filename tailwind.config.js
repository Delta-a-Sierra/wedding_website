/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.templ", "./**/.html"],
  theme: {
    extend: {
      colors: {
        black: "#0d0d0b",
        white: "#ffffff",
        primary: "#355e3b",
        "off-grey": "#a7a096",
        "grey-txt": "#656565",
      },
      height: {
        200: "50rem",
      },
      fontFamily: {
        italiana: ["Italiana", "sans-serif"],
        roboto: ["Roboto", "serif"],
      },
    },
  },
  plugins: [],
};
