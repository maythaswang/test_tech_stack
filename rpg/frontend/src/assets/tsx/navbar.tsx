import React from 'react';
import '../css/navbar.css';

const Navbar: React.FC = () => {
  return (
    <nav className="navbar">
      <ul>
        <li><a href="/">Home</a></li>
        <li><a href="/blogs">Blogs</a></li>
      </ul>
    </nav>
  );
};

export default Navbar;