module.exports = {
  future: {
    removeDeprecatedGapUtilities: true,
    purgeLayersByDefault: true,
  },
  purge: ['./src/components/**/*.{js,ts,jsx,tsx}', './src/pages/**/*.{js,ts,jsx,tsx}'],
  theme: {
    extend: {
      // colors: {
      //   'accent-1': '#333',
      // },
      colors: {
        'accent-1': '#FAFAFA',
        'accent-2': '#EAEAEA',
        'accent-7': '#333',
        success: '#0070f3',
        cyan: '#79FFE1',
      },
      spacing: {
        28: '7rem',
      },
      letterSpacing: {
        tighter: '-.04em',
      },
      lineHeight: {
        tight: 1.2,
      },
      fontSize: {
        '5xl': '2.5rem',
        '6xl': '2.75rem',
        '7xl': '4.5rem',
        '8xl': '6.25rem',
      },
      boxShadow: {
        small: '0 5px 10px rgba(0, 0, 0, 0.12)',
        medium: '0 8px 30px rgba(0, 0, 0, 0.12)',
      },

    },
  },
  variants: {},
  plugins: [
    require('@tailwindcss/custom-forms'),
  ],
}

// module.exports = {
//   purge: ['./src/components/**/*.tsx', './src/pages/**/*.tsx'],
//   theme: {
//     extend: {
//       colors: {
//         'accent-1': '#FAFAFA',
//         'accent-2': '#EAEAEA',
//         'accent-7': '#333',
//         success: '#0070f3',
//         cyan: '#79FFE1',
//       },
//       spacing: {
//         28: '7rem',
//       },
//       letterSpacing: {
//         tighter: '-.04em',
//       },
//       lineHeight: {
//         tight: 1.2,
//       },
//       fontSize: {
//         '5xl': '2.5rem',
//         '6xl': '2.75rem',
//         '7xl': '4.5rem',
//         '8xl': '6.25rem',
//       },
//       boxShadow: {
//         small: '0 5px 10px rgba(0, 0, 0, 0.12)',
//         medium: '0 8px 30px rgba(0, 0, 0, 0.12)',
//       },
//     },
//   },
// }