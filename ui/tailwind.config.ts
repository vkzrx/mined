import type { Config } from "tailwindcss";
import plugin from "tailwindcss/plugin";

const config: Config = {
  content: [
    "./src/pages/**/*.{ts,tsx}",
    "./src/components/**/*.{ts,tsx}",
    "./src/app/**/*.{ts,tsx}",
  ],
  theme: {
    extend: {
      animation: {
        spinner: "spinner 1.2s cubic-bezier(0.4, 0.5, 0.5, 1) infinite",
      },
      keyframes: {
        spinner: {
          "0%": { top: "0px", height: "16px" },
          "50%, 100%": { top: "16px", height: "8px" },
        },
      },
    },
  },
  plugins: [
    plugin(({ matchUtilities, theme }) => {
      matchUtilities(
        {
          "animation-delay": (value) => {
            return {
              "animation-delay": value,
            };
          },
        },
        {
          values: theme("transitionDelay"),
        },
      );
    }),
  ],
};

export default config;
