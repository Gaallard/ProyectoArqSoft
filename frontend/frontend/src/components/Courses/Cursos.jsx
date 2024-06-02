import React, { useState } from 'react';
import './Cursos.css';

const Cursos = () => {
  // Datos de los cursos
  const cursosIniciales = [
    { id: 1, nombre: 'React Básico', descripcion: 'Introducción a React.js y sus conceptos básicos.' },
    { id: 2, nombre: 'JavaScript Avanzado', descripcion: 'Profundización en JavaScript moderno y sus características avanzadas.' },
    { id: 3, nombre: 'Node.js', descripcion: 'Desarrollo de aplicaciones del lado del servidor con Node.js.' },
    { id: 4, nombre: 'TypeScript', descripcion: 'Aprende a utilizar TypeScript para un desarrollo más seguro.' },
    { id: 5, nombre: 'GraphQL', descripcion: 'Consultas eficientes con GraphQL.' },
    { id: 6, nombre: 'Redux', descripcion: 'Gestión de estado avanzada con Redux.' },
    { id: 7, nombre: 'React Native', descripcion: 'Desarrollo de aplicaciones móviles con React Native.' },
    { id: 8, nombre: 'Next.js', descripcion: 'Renderizado del lado del servidor y generación estática con Next.js.' },
    { id: 9, nombre: 'Webpack', descripcion: 'Empaquetado de módulos con Webpack.' },
  ];

  // Estado para el término de búsqueda
  const [busqueda, setBusqueda] = useState('');
  const [cursos, setCursos] = useState(cursosIniciales);

  // Manejar el cambio en el campo de búsqueda
  const manejarCambio = (e) => {
    setBusqueda(e.target.value);
    filtrarCursos(e.target.value);
  };

  // Filtrar cursos basado en la búsqueda
  const filtrarCursos = (termino) => {
    const cursosFiltrados = cursosIniciales.filter((curso) =>
      curso.nombre.toLowerCase().includes(termino.toLowerCase())
    );
    setCursos(cursosFiltrados);
  };

  // Manejar la inscripción al curso
  const manejarInscripcion = (nombreCurso) => {
    alert(`Te has inscrito en el curso: ${nombreCurso}`);
  };

  return (
    <div className="container">
      <h1>Cursos de Programación</h1>
      <input
        type="text" className='busqueda'
        placeholder="Buscar curso..."
        value={busqueda}
        onChange={manejarCambio}
      />
      <ul className="grid">
        {cursos.map((curso) => (
          <li key={curso.id} className="card">
            <h2>{curso.nombre}</h2>
            <p>{curso.descripcion}</p>
            <button onClick={() => manejarInscripcion(curso.nombre)}>
              Inscribirse
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Cursos;
