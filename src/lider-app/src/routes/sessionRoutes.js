import { lazy } from 'react';

const Home = lazy(() => import('../pages/Home'));
const Catalog = lazy(() => import('../pages/Catalog'));

const sessionRoutes = {
    Home: {
        component: Home,
        path: '/'
    },
    Catalog: {
        component: Catalog,
        path: '/catalog'
    }
};

export default sessionRoutes;