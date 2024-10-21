import React from 'react';
import './ToggleSwitchUI.css';

const ToggleSwitchUI = ({ isOn, handleToggle, onColor, offColor, isLoading }) => {
    return (
        <div
            className={`toggle-switch-ui ${isOn ? 'active' : ''}`}
            onClick={isLoading ? null : handleToggle} // Отключаем нажатие, если идет загрузка
            style={{ backgroundColor: isLoading || isOn ? onColor : offColor }}
        >
            {isLoading ? <div className="loading-lines"></div> : null}
            <div className={`toggle-circle-ui ${isOn || isLoading ? 'circle-active' : ''}`}></div>
        </div>
    );
};

export default ToggleSwitchUI;
