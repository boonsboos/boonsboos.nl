/** @type {import('tailwindcss').Config} */
module.exports = {
  safelist: [
    'chroma',
    'line', 'k', 'kc', 'kd', 'kn', 'kr', 'kt', 'nt', 'kp', 'nx', 'c',
    's','sr', 'nf'
  ],
  content: ["./html/**/*.html"],
  theme: {
    extend: {},
    fontFamily: {
      mono: ['Roboto Mono', 'monospace'],
      sans: ['Roboto Mono', 'monospace']
    },
  },
  plugins: [
    require('@tailwindcss/container-queries'),
    require('@tailwindcss/typography')
  ],
}

