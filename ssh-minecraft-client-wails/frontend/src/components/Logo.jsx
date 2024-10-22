import React from 'react';
import './Logo.css'; // Импортируем CSS для стилизации

const Logo = () => {
    return (
        <h1 className="neon" data-text="U">
            СКБ <span className="flicker-slow">Б</span>
            Р
            <span className="flicker-fast">А</span>
            З
            <span className="flicker-fast">И</span>
            ЛИЯ
        </h1>
    );
};

export default Logo;
