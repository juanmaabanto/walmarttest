// import { productApiUrl } from '../config';
import axios from 'axios';
import React, { useReducer, useRef } from 'react';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Checkbox from '@mui/material/Checkbox';
import FormControlLabel from '@mui/material/FormControlLabel';
import Link from '@mui/material/Link';
import TextField from '@mui/material/TextField';
import useMediaQuery from '@mui/material/useMediaQuery';

export const initialState = {
    error: '',
    search: '',
    porId: false,
    data: [],
    total: 0,
    pageSize: 10,
    start: 0
};

export const reducer = (state, action) => {
    switch (action.type) {
        case 'error':
            return {
                ...state,
                error: action.error
            };
        case 'field':
            return {
                ...state,
                [action.key]: action.value
            };
        case 'fetch-finish':
            return {
                ...state,
                data: action.data,
                total: action.total,
                pageSize: action.pageSize,
                start: action.start
            }
        default:
            return state;
    }
};

const Catalog = ({ route }) => {
    const matches = useMediaQuery('(min-width:600px)');
    const [state, dispatch] = useReducer(reducer, initialState);
    const inputEl = useRef(null);

    const onFieldChange = event => {
        dispatch({ type: 'field', key: event.target.name, value: event.target.value });
    };

    const fetchProducts = async () => {
        if(state.search.trim() === '' || state.search.trim().length < 3 ) {
            dispatch({ type: 'error', error: 'Ingreso al menos 3 caracteres' });
            return;
        }

        dispatch({ type: 'error', error: '' });

        try {
            let productApiUrl = 'http://localhost:5500/';
            let url = state.porId ? `${productApiUrl}api/v1/products/${state.search}`
                : `${productApiUrl}api/v1/products/?search=${state.search}`;

            let response = await axios.get(url);
            let data = response.data;

            dispatch({ type: 'fetch-finish', data: data.data, total: data.total, pageSize: data.pageSize, start: data.start });

            console.log(state);
        } catch (error) {
            inputEl.current.focus();

            if(error.response) {
                dispatch({ type: 'error', error: error.response.data.message });
            }
            else {
                dispatch({ type: 'error', error: 'Se produjo un error al conectarse al servidor.' })
            }
        }
    };

    return (
        <Box>
        <Box sx={{
            display: 'flex',
            margin: '48px 64px 0'
          }}
        >
            <Link
                component="button"
                variant="body2"
                onClick={() => route.history.push({
                    pathname: '/',
                    search: ''
                })}
            >
                Regresar
            </Link>
            <Box sx={{
                flex: 1,
                textAlign: 'center',
                fontSize: 36,
                fontWeight: 600
            }}>
                Catálogo de Productos
            </Box>
        </Box>
        <Box>
            <Box sx={{ padding: '2rem 0 0' }}>
                <Box
                    sx={{
                        alignItems: 'center',
                        display: 'flex',
                        flexDirection: 'column',
                        mt: '1.5rem',
                    }}
                >
                    <TextField
                        autoFocus
                        error={state.error !== ''}
                        helperText={state.error}
                        inputRef={inputEl}
                        label='Buscar Productos'
                        name='search'
                        placeholder='Buscar Productos'
                        value={state.search}
                        variant={matches ? 'outlined' : 'standard'}
                        onChange={onFieldChange}
                        onKeyPress={(event) => {if(event.key === 'Enter') fetchProducts()}}
                        inputProps={{
                            maxLength: 50
                        }}
                    />

                        <Box sx={{ color: 'text.primary' }}>
                            <FormControlLabel
                                classes={{
                                    label: { fontsize: '0.575rem' }
                                }}
                                control={<Checkbox
                                    color="primary"
                                    sx={{ p: '0.375rem' }}
                                    value={state.porId}
                                    onChange={event => dispatch({ type: 'field', key: 'porId', value: event.target.checked })}
                                />}
                                label="Buscar Por Id"
                            />
                        </Box>
                        <Button color="primary" variant="contained" onClick={() => fetchProducts()}>
                            Buscar
                        </Button>
                    </Box>
            </Box>
        </Box>
        </Box>
    );
};

export default Catalog;