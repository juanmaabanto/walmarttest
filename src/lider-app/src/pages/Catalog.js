// import { productApiUrl } from '../config';
import axios from 'axios';
import React, { useEffect, useReducer, useRef } from 'react';
import Badge from '@mui/material/Badge';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Checkbox from '@mui/material/Checkbox';
import CircularProgress from '@mui/material/CircularProgress';
import FormControlLabel from '@mui/material/FormControlLabel';
import Link from '@mui/material/Link';
import Pagination from '@mui/material/Pagination';
import Stack from '@mui/material/Stack';
import TextField from '@mui/material/TextField';
import useMediaQuery from '@mui/material/useMediaQuery';

export const initialState = {
    error: '',
    search: '',
    porId: false,
    data: [],
    total: -1,
    pageSize: 10,
    start: 0,
    isLoading: false
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
        case 'change-page':
            return {
                ...state,
                start: action.start
            };
        case 'fetch-start':
            return {
                ...state,
                data: [],
                isLoading: true,
                total: 0
            };
        case 'fetch-error':
            return {
                ...state,
                data: [],
                isLoading: false,
                total: 0,
                error: action.error
            };
        case 'fetch-finish':
            return {
                ...state,
                data: action.data,
                isLoading: false,
                total: action.total,
                pageSize: action.pageSize,
                start: action.start
            };
        default:
            return state;
    }
};

const Catalog = ({ route }) => {
    const matches = useMediaQuery('(min-width:600px)');
    const [state, dispatch] = useReducer(reducer, initialState);
    const inputEl = useRef(null);

    useEffect(() => {
        if(state.total >= 0) {
            fetchProducts();
        }
    // eslint-disable-next-line
    }, [state.start])

    const onFieldChange = event => {
        dispatch({ type: 'field', key: event.target.name, value: event.target.value });
    };

    const handlePageChange = (event, value) => {
        console.log(parseInt((value - 1)*state.pageSize))
        dispatch({ type: 'change-page', start: parseInt((value - 1)*state.pageSize) });
    };

    const fetchProducts = async () => {
        if (state.isLoading) {
            return;
        }
        if(state.search.trim() === '' || (state.search.trim().length < 3 && !state.porId) ) {
            dispatch({ type: 'error', error: 'Ingreso al menos 3 caracteres' });
            return;
        }

        dispatch({ type: 'error', error: '' });
        dispatch({ type: 'fetch-start' })

        try {
            let porId = state.porId;
            let productApiUrl = 'http://localhost:5500/';
            let url = porId ? `${productApiUrl}api/v1/products/${state.search}`
                : `${productApiUrl}api/v1/products?search=${state.search}&pageSize=${state.pageSize}&start=${state.start}`;

            let response = await axios.get(url);
            let data = response.data;

            if (porId) {
                var list = []

                list.push(data);
                dispatch({ type: 'fetch-finish', data: list, total: 1, pageSize: 10, start: 0 });
            } else {
                dispatch({ type: 'fetch-finish', data: data.data, total: data.total, pageSize: data.pageSize, start: data.start });
            }
        } catch (error) {
            inputEl.current.focus();

            if(error.response) {
                dispatch({ type: 'fetch-error', error: error.response.data.message });
            }
            else {
                dispatch({ type: 'fetch-error', error: 'Se produjo un error al conectarse al servidor.' });
            }
        }
    };

    return (
        <Box>
        <Box sx={{
            display: 'flex',
            margin: '36px 64px 0'
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
            <Box sx={{ padding: '1rem 0 0' }}>
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
        {state.isLoading && (<Box sx={{ alignItems: 'center', display: 'flex', flexDirection: 'column', m: '36px' }}>
            <CircularProgress />
        </Box>)}
        {!state.isLoading && (<Box component="ul" cols={5} sx={{ alignItems: 'center', m: '36px' }}>
        {state.data.length === 0 && (<Box sx={{textAlign: 'center'}}>No hay información que mostrar</Box>)}
        {state.data.length > 0 && state.data.map((item, index) => (
            <Box
                component="li"
                key={index}
                sx={{
                    display: 'inline-block',
                    height: '20rem',
                    listStyleType: 'none',
                    padding: '0.5rem',
                    textAlign: 'center',
                    fontsize: '10px',
                    width: '10rem'
                }}
            >
                <Box component="img" src={"https://" + item.image} alt={item.description} />
                <Box><Box component="span" sx={{ fontWeight: 600 }}>{item.brand}</Box> {item.description}</Box>
                <Badge
                    anchorOrigin={{
                        vertical: 'bottom',
                        horizontal: 'right',
                    }}
                    badgeContent={item.conDescuento ? '50%' : ''}
                    color="secondary"
                    invisible={!item.conDescuento}
                >
                    <Box sx={{ paddingRight: '24px' }}>${item.price}</Box>
                </Badge>
            </Box>
        ))}
        </Box>)}

        <Box>
            <Stack spacing={2} sx={{
                alignItems: 'center',
                marginTop: '2rem'
            }}>
                <Pagination
                    count={parseInt(state.total / state.pageSize) + 1}
                    page={parseInt(state.start / state.pageSize) + 1}
                    onChange={handlePageChange}
                    color="primary"
                />
            </Stack>
        </Box>
        </Box>
    );
};

export default Catalog;