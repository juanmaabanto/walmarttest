import React, { Suspense } from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Container from '@mui/material/Container';
import Toolbar from '@mui/material/Toolbar';

const Main = ({component, route}) => {
    const Component = component;

    return (
        <Box
            sx={{
                bgcolor: '#fff',
                minHeight: '100%',
                position: 'relative'
            }}
        >
            <AppBar position="static">
                <Container maxWidth="false">
                    <Toolbar disableGutters>
                        <Box
                            noWrap
                            component="img"
                            sx={{ height: 40 }}
                            src="https://www.lider.cl/images/logo.svg"
                        />
                    </Toolbar>
                </Container>
            </AppBar>
            <Suspense fallback={<div></div>}>
                <Component
                    route={route}
                />
            </Suspense>
        </Box>
    );
};

export default Main;