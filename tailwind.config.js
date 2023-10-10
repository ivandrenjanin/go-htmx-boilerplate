/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./internal/templates/**/*.tmpl"], // This is where your HTML templates / JSX files are located
  theme: {},
  plugins: [
    require('@tailwindcss/typography'),
    require('@tailwindcss/forms'),
    require('@tailwindcss/aspect-ratio'),
    require('@tailwindcss/container-queries'),
  ],
};
