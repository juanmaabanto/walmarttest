import React, { Suspense } from 'react';
import Box from '@mui/material/Box';

const Main = ({component, route}) => {
    const Component = component;

    return (
        <Box
            sx={{
                bgcolor: { sm: 'background.default' },
                display: 'flex',
                flexDirection: { sm: 'column' },
                minHeight: '100%',
                position: { sm: 'relative' }
            }}
        >
            <Suspense fallback={<div></div>}>
                <Component
                    route={route}
                />
            </Suspense>
        </Box>
    );
};

export default Main;