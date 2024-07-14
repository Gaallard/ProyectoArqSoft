import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import './Cursos.css';
import Swal from 'sweetalert2';

const Cursos = () => {
  const navigate = useNavigate();
  const [busqueda, setBusqueda] = useState('');
  const [cursos, setCursos] = useState([]);
  const [cursosFiltrados, setCursosFiltrados] = useState([]);

  useEffect(() => {
    const token = sessionStorage.getItem('token');
    if (!token) {
      navigate('/login');
    } else {
      cargarCursos(); // Cargar cursos si el usuario está autenticado
    }
  }, [navigate]);

  const cargarCursos = async () => {
    try {
      const token = sessionStorage.getItem('token');
      const response = await fetch('http://localhost:8080/courses', {
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${token}`
        }
      });
  
      if (!response.ok) {
        throw new Error('Error en la carga de cursos');
      }
  
      const data = await response.json();
      setCursos(data); // Asignar los cursos obtenidos al estado cursos
      setCursosFiltrados(data);
    } catch (error) {
      console.error("Error durante la carga de cursos:", error);
      alert('Error durante la carga de cursos');
    }
  };

  const inscripcion = async (curso) => {
    try {
      const token = sessionStorage.getItem('token');
      const idCourse = curso.id; // Suponiendo que el objeto curso tiene un atributo id que representa el ID del curso
      const url = `http://localhost:8080/inscriptions/${idCourse}`;
      const response = await fetch(url, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
        }
      });
      if (response.status === 201) {
        Swal.fire('Inscripción exitosa', 'Bienvenido al curso', 'success');
      } else {
        Swal.fire('Error', 'No se pudo inscribir al curso', 'error');
      }
    } catch (error) {
      console.error("Error durante la inscripción:", error);
      alert('Error durante la inscripción');
    }
  };

  const manejarCambio = (e) => {
    const value = e.target.value.toLowerCase();
    setBusqueda(value);
    const nuevosCursos = cursos.filter((curso) => {
      return curso.name && curso.name.toLowerCase().includes(value);
    });
    setCursosFiltrados(nuevosCursos);
  };

  return (
    <div className="container">
      <h1>Cursos de Programación</h1>
      <input
        type="text"
        className='busqueda'
        placeholder="Buscar curso..."
        onChange={manejarCambio}
        value={busqueda}
      />
      <ul className="grid">
        {cursosFiltrados.map((curso) => (
          <li key={curso.id} className="card">
            <h2>{curso.name}</h2>
            <p>{curso.description}</p>
            <button onClick={() => inscripcion(curso)}>
              Inscribirse
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Cursos;
