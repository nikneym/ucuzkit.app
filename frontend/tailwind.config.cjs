/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        zilla: "Zilla Slab, serif",
        itim: "Itim, cursive",
      },

      boxShadow: {
        "inner-xl": "inset 0px 0px 20px 1px rgba(36,36,36,1)",
      },
    },
  },
  plugins: [],
}
