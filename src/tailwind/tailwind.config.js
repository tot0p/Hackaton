/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/**/*.{html,js}",
  ],
  theme: {
    extend: {
      colors: {
        'background': '#FFFFFF',
        'primary': '#6892D5',
        'secondary': '#2A3B57',
        'tertiary': '#0B84D9',
      },
      fontFamily: {
        'primary': ['IntegralCF'],
        'secondary': ['CriteriaCF'],
        'tertiary': ['SpaceGrotesk'],
      }
    },
  },
  plugins: [],
}
