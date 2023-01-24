/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
    backgroundColor: theme => ({
      ...theme('colors'),
      'primary': '#FFEFD5',
      'secondary': '#ffed4a',
      'danger': '#e3342f',
    })
  },
  plugins: [],
}