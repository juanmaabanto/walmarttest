import React, { lazy, Suspense } from 'react';
import { Route, Switch } from 'react-router-dom';
import { MyTheme } from './helpers/theme';
import ThemeProvider from '@mui/material/styles/ThemeProvider';

import Main from './layouts/Main';
import sessionRoutes from './routes/sessionRoutes';

const NotFound = lazy(() => import('./pages/NotFound'));

const App = () => {
    const theme = MyTheme();

    return(
        <ThemeProvider theme={theme}>
            <Suspense fallback={<div></div>}>
                <Switch>
                    {Object.keys(sessionRoutes).map((route, key) => {
                        const path = sessionRoutes[route].path;
                        const component = sessionRoutes[route].component;

                        return (
                            <Route
                                exact
                                key={key}
                                path={path}
                                render={(r) => 
                                    <Main
                                        component={component}
                                        route={r}
                                    />
                                }
                            />
                        );
                    })}
                    <Route component={NotFound} />
                </Switch>
            </Suspense>
        </ThemeProvider>
    );
};

export default App;