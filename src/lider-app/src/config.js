import { config } from 'dotenv';

config({
    path: `../.env.${process.env.NODE_ENV}`
});

export const productApiUrl = process.env.REACT_APP_PRODUCT_API_URL;