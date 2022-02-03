import React from 'react';
import Box from '@mui/material/Box';

// xs, extra-small: 0px
// sm, small: 600px
// md, medium: 900px
// lg, large: 1200px
// xl, extra-large: 1536px

const NotFound = () => {
    return (
        <Box sx={{
            fontSize: { xs: 10, sm: 20, md: 30, lg: 40, xl: 50 }
          }}>
            NotFound
        </Box>
    );
};

export default NotFound;