import createTheme from '@mui/material/styles/createTheme';

const theme = createTheme({
    components: {
        MuiButton: {
            styleOverrides: {
                root: {
                    fontWeight: 500,
                    textTransform: 'none'
                }
            }
        },
        MuiChip: {
            styleOverrides: {
                root: {
                    fontSize: '1rem'
                }
            }
        },
        MuiFormControlLabel: {
            styleOverrides: {
                label: {
                    fontsize: '0.875rem'
                }
            }
        },
        MuiFormHelperText: {
            styleOverrides: {
                root: {
                    fontSize: '0.8125rem'
                }
            }
        },
        MuiOutlinedInput: {
            styleOverrides: {
                adornedEnd: {
                    paddingRight: 0
                }
            }
        }
    },
    typography: {
        fontFamily: [
            'Montserrat',
            'sans-serif'
        ].join(',')
    }
});

const themes = {
    default: createTheme({
        palette: {
            primary: {
                main: '#0071dc'
            }
        },
        components: theme.components,
        typography: theme.typography
    })
};

export const MyTheme = () => {
    return themes.default;
}