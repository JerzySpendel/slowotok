import React from 'react';

export const Header = ({ startGame }) => (
    <header className="header">
        <div className="header-wrapper">        
            <p className="logo">Wordament</p>
            <button className="btn" onClick={() => startGame()}> Start game </button>
        </div>
    </header>
)