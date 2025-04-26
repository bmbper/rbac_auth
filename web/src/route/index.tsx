import { RouteObject } from 'react-router-dom';
import HomePage from '../page/home';
import LoginPage from '../page/login';

export const routes: RouteObject[] = [
  {
    path: '/',
    element: <HomePage />,
  },
  {
    path: '/login',
    element: <LoginPage />,
  },
];