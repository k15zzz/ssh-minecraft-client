import React from 'react';
import './ButtonUI.css'; // Импортируем CSS для стилизации

const ButtonUI = ({ label, onClick, isConnected, ...props }) => {
    // Определяем стиль кнопки в зависимости от состояния подключения
    const buttonClass = isConnected ? 'button-ui disconnect' : 'button-ui connect';

    return (
        <button className={buttonClass} onClick={onClick} {...props}>
            {label}
        </button>
    );
};

export default ButtonUI;
