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
        'secondary': '#C9FDD7',
        'tertiary': '#84CAF0',
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
