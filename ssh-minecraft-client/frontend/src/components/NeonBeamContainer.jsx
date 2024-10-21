import React from 'react';
import './NeonBeamContainer.css';

const NeonBeamContainer = ({ children, status }) => {
    return (
        <div className={`neon-container ${status ? 'active' : ''}`}>
            {/*<div className="neon-glow"></div>*/}
            <div className="content">
                {children}
            </div>
        </div>
    );
};

export default NeonBeamContainer;
