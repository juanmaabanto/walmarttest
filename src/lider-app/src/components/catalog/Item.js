import React from 'react';
import Badge from '@mui/material/Badge';
import Box from '@mui/material/Box';

const Item = ({ data }) => {

    return (
        <Box
            component="li"
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
            <Box component="img" src={"https://" + data.image} alt={data.description} />
            <Box>
                <Box component="span" sx={{ fontWeight: 600 }}>{data.brand}</Box> {data.description}
            </Box>
            <Badge
                anchorOrigin={{
                    vertical: 'bottom',
                    horizontal: 'right',
                }}
                badgeContent={data.conDescuento ? '50%' : ''}
                color="secondary"
                invisible={!data.conDescuento}
            >
                <Box sx={{ paddingRight: '24px' }}>${data.price}</Box>
            </Badge>
        </Box>
    );
};

export default Item;