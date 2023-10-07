/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './index.html',
    './src/**/*.{vue,js,ts,jsx,tsx}',
  ],
  theme: {
    screens: {
      '2xs': '320px',
      'xs': '360px',
      'sm': '414px',
      'md': '712px',
      '2md': '768px',
      'lg': '1018px',
      '2lg': '1100px',
      'xl': '1358px',
      '2xl': '1536px'
    },

    colors: {
      transparent: 'rgba(0, 0, 0, 0)',
      current: 'currentColor',


      primary: {
        brand: {
          accent: '#FF5132',
          onAccent: '#FF2600',

          primary: '#2B2D33',
          onPrimary: '#000000',

          wildSand: '#F5F5F7',

          white: '#FFFFFF',
        },

        accentOpacity: {
          3: 'rgba(255, 81, 50, 0.03)',
          7: 'rgba(255, 81, 50, 0.07)',
          14: 'rgba(255, 81, 50, 0.14)',
        },

        atlantic: {
          10: '#E6F3F9',
          30: '#C2E1F0',
          60: '#66B5D7',
          100: '#339DC9',
          120: '#0084BC',
          140: '#00638D',
          160: '#045577',
        },

        seattle: {
          1: '#FBFBFB',
          5: '#F8F8F8',
          10: '#E9EAEA',
          30: '#D5D5D6',
          60: '#AAABAD',
          100: '#808185',
          120: '#55575C',
          140: '#404247',
        },

        london: {
          40: 'rgba(255, 255, 255, 0.64)',
          100: 'rgba(32, 35, 51, 0.07)',
          110: 'rgba(32, 35, 51, 0.14)',
          120: 'rgba(32, 35, 51, 0.28)',
        },
      },

      utilitarian: {
        geneva: {
          10: '#E8F7F0',
          30: '#BFE9D4',
          60: '#52C58B',
          100: '#0DC268',
          120: '#008E47',
          200: '#1A7045'
        },

        moscow: {
          10: '#FEEAE7',
          30: '#FBD5CE',
          60: '#F4806C',
          100: '#ED2C0A',
          120: '#B22108',
        },

        osaka: {
          10: '#FFF6E6',
          30: '#FFECCC',
          60: '#FFC566',
          100: '#FF9E00',
          120: '#BF7700',
        },
      },

      secondary: {
        basic: {
          miami: {
            10: '#E7F8F4',
            30: '#CFF1E9',
            100: '#5ED1B6',
          },

          sochi: {
            10: '#E6F4FC',
            30: '#CDE9FA',
            100: '#57B6ED',
          },

          paris: {
            10: '#F0F0FE',
            30: '#E1E0FE',
            100: '#9A99FB',
            200: '#3332BB'
          },

          tokyo: {
            10: '#FAEFFA',
            30: '#F5DFF5',
            100: '#DC94DD',
          },

          dubai: {
            10: '#FFF4E8',
            30: '#FEE9D1',
            100: '#FCB664',
          },

          nice: {
            10: '#FFEFEF',
            30: '#FEE0E0',
            100: '#FC9696',
          },

          amsterdam: {
            100: '#D917A9',
          },

          petersburg: {
            10: '#F2F1EE',
            30: '#E3E2DF',
            100: '#73726D',
          },
        },
      },
    },

    boxShadow: {
      'elevation-0': '0px 0px 0px 1px rgba(0, 44, 94, 0.1)',
      'elevation-1': '0px 2px 4px 0px #20233305, 0px 1px 1px 0px #2023330A',
      'elevation-2': '0px 2px 4px -1px #20233326, 0px 0px 0px 1px #2023331A',
      'elevation-3': '0px 8px 20px 0px #070F3C1A, 0px 0px 1px 0px #070F3C66',
      'elevation-4': '0px 1px 2px rgba(0, 44, 94, 0.12), 0px 0px 2px rgba(0, 44, 94, 0.15), 0px 14px 28px -14px rgba(0, 44, 94, 0.2)',
      'elevation-5': '0px 1px 2px rgba(0, 44, 94, 0.12), 0px 0px 2px rgba(0, 44, 94, 0.15), 0px 21px 42px -21px rgba(0, 44, 94, 0.2)',
      'elevation-6': '0px 1px 2px 0px #0000001F, 0px 8px 24px 0px #0000001F',
      'none': 'none',
    },

    fontFamily: {
      sans: ['GolosTextVF', 'sans-serif'],
    },

    extend: {
      fontSize: {
        'h1': ['36px', {
          lineHeight: '40px',
          fontWeight: '700',
        }],

        'h2': ['32px', {
          lineHeight: '40px',
          fontWeight: '700',
        }],

        'h3': ['28px', {
          lineHeight: '32px',
          fontWeight: '700',
        }],

        'h4': ['24px', {
          lineHeight: '30px',
          fontWeight: '700',
        }],

        'h5': ['20px', {
          lineHeight: '24px',
          fontWeight: '700',
        }],

        'h6': ['18px', {
          lineHeight: '22px',
          fontWeight: '700',
        }],

        'xl-18': ['18px', {
          lineHeight: '28px',
        }],

        'l-long-16': ['16px', {
          lineHeight: '24px',
        }],

        'l-short-16': ['16px', {
          lineHeight: '20px',
        }],

        'm-14': ['14px', {
          lineHeight: '20px',
        }],

        's-12': ['12px', {
          lineHeight: '16px',
        }],

        'caption-14': ['14px', {
          lineHeight: '16px',
          letterSpacing: '0.01em',
        }],

        'caption-12': ['12px', {
          lineHeight: '16px',
          letterSpacing: '0.01em',
        }],
      },
    },
  },
  plugins: [],
}

