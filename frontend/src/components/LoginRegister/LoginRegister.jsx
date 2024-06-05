import React, { useState } from "react";
import './LoginRegister.css';
import { FaUserAlt, FaLock, FaEnvelope } from "react-icons/fa";
import { Link } from "react-router-dom";

const LoginRegister = () => {
  const [action, setAction] = useState('');
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [registerUsername, setRegisterUsername] = useState('');
  const [email, setEmail] = useState('');
  const [registerPassword, setRegisterPassword] = useState('');

  const registerLink = () => {
    setAction('active');
  };

  const loginLink = () => {
    setAction('');
  };

  const isLoginComplete = username && password;
  const isRegisterComplete = registerUsername && email && registerPassword;



  return (
    <div className={`wrapper ${action}`}>
      <div className="from-box login">
        <form action="">
          <h1>Login</h1>
          <div className="input-box">
            <input 
              type="text" 
              placeholder="Usuario" 
              required 
              value={username} 
              onChange={(e) => setUsername(e.target.value)} 
            />
            <FaUserAlt className="icon" />
          </div>
          <div className="input-box">
            <input 
              type="password" 
              placeholder="Contraseña" 
              required 
              value={password} 
              onChange={(e) => setPassword(e.target.value)} 
            />
            <FaLock className="icon" />
          </div>
          <div className="olvidar">
            <a href="#">¿Olvidaste la contraseña?</a>
          </div>
          <Link to='/Cursos'>
            <button type="submit" disabled={!isLoginComplete}>Login</button>
          </Link>
          <div className="register">
            <p>No tienes una cuenta? <a href="#" onClick={registerLink}>Regístrate</a></p>
          </div>
        </form>
      </div>

      <div className="from-box register">
        <form action="">
          <h1>Regístrate</h1>
          <div className="input-box">
            <input 
              type="text" 
              placeholder="Usuario" 
              required 
              value={registerUsername} 
              onChange={(e) => setRegisterUsername(e.target.value)} 
            />
            <FaUserAlt className="icon" />
          </div>
          <div className="input-box">
            <input 
              type="email" 
              placeholder="Email" 
              required 
              value={email} 
              onChange={(e) => setEmail(e.target.value)} 
            />
            <FaEnvelope className="icon" />
          </div>
          <div className="input-box">
            <input 
              type="password" 
              placeholder="Contraseña" 
              required 
              value={registerPassword} 
              onChange={(e) => setRegisterPassword(e.target.value)} 
            />
            <FaLock className="icon" />
          </div>
          <button type="submit" disabled={!isRegisterComplete}>Registrarse</button>
          <div className="register-link">
            <p>¿Ya tienes una cuenta? <a href="#" onClick={loginLink}>Login</a></p>
          </div>
        </form>
      </div>
    </div>
  );
}

export default LoginRegister;
