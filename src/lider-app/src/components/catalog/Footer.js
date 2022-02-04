import React from 'react';
import Box from '@mui/material/Box';
import Pagination from '@mui/material/Pagination';
import Stack from '@mui/material/Stack';

const Footer = ({ pageSize, start, total, handlePageChange }) => {

    return (
        <Box>
            <Stack spacing={2} sx={{
                alignItems: 'center',
                marginTop: '2rem'
            }}>
                <Pagination
                    count={parseInt(total / pageSize) + 1}
                    page={parseInt(start / pageSize) + 1}
                    onChange={handlePageChange}
                    color="primary"
                />
            </Stack>
        </Box>
    );
};

export default Footer;