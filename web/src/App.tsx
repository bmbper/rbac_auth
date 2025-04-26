import '@arco-design/web-react/dist/css/arco.css';
import { useRoutes } from 'react-router-dom';
import './App.scss';
import { routes } from './route';

const App = () => {
  const element = useRoutes(routes);
  return <div className="app-root">{element}</div>;
};

export default App;
