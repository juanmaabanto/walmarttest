import React from 'react';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';

const Home = ({ route }) => {
    return (
        <Box sx={{
            alignItems: 'center',
            display: 'flex',
            flexDirection: 'column',
            fontSize: { xs: 10, sm: 20, md: 30, lg: 40, xl: 50 },
            margin: '100px 0',
            minHeight: '100%'
          }}>
            <Box sx={{
                fontSize: 36,
                fontWeight: 600
            }}>
                Bienvenido
            </Box>
            <Box sx={{
                padding: 2
            }}>
                <Button
                    variant="contained"
                    onClick={() => route.history.push({
                        pathname: 'catalog',
                        search: route.location.search
                    })}
                >
                    Ver Catálogo
                </Button>
            </Box>
        </Box>
    );
};

export default Home;