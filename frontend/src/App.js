import './App.css';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { Main } from './components/desktop/pages/main/Main';
import { Support } from './components/desktop/pages/support/Support';
import { Profile } from './components/desktop/pages/profile/Profile';
import { Login } from './components/desktop/pages/login/Login'; 
import { Hotels } from './components/desktop/pages/hotels/Hotels';
import { Order } from './components/desktop/pages/order/Order'; 
import { Favorite } from './components/desktop/pages/favorite/Favorite';
import { Registration } from './components/desktop/pages/registration/Registration';

const router = createBrowserRouter([
  {
    path: "/",
    element: <Main />
  }, 
  {
    path: "/support",
    element: <Support />
  },
  {
    path: "/favorite",
    element: <Favorite />
  },
  {
    path: "/profile",
    element: <Profile />
  },
  {
    path: "/login",
    element: <Login />
  },
  {
    path: "/registration",
    element: <Registration />
  },
  {
    path: "/hotels",
    element: <Hotels />
  },
  {
    path: "/order",
    element: <Order />
  }
])

export function App() {
  return (
    <RouterProvider router={router} />
  );
}
