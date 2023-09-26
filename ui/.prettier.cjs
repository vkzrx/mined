/** @type {import("prettier").Config} */
const config = {
  printWidth: 80,
  tabWidth: 2,
  semi: true,
  trailingComma: "all",
  singleQuote: true,
  jsxSingleQuote: false,
  quoteProps: "as-needed",
  arrowParens: "always",
  bracketSpacing: true,
  bracketSameLine: false,
  plugins: ["prettier-plugin-tailwindcss"],
};

module.exports = config;
