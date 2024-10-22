import React from 'react';
import './ErrorContainer.css';

const ErrorContainer = ({ error }) => {
    // Контейнер отображается только если есть ошибка
    if (!error) return null;

    return (
        <div className="error-container">
            <p className="error-message">{error}</p>
        </div>
    );
};

export default ErrorContainer;
