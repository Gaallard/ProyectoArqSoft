import './App.css';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import LoginRegister from './components/LoginRegister/LoginRegister.jsx';
import Cursos from './components/Courses/Cursos.jsx';

function App() {
  return (
    <BrowserRouter>
    <div className="App">
      <Routes>
      <Route path="/users" element={<LoginRegister />} />
      <Route path="/courses" element={<Cursos />} />
      </Routes>
    </div>
    </BrowserRouter>
  );
}

export default App;
