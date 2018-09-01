import React from 'react';
import logo from '../logo.svg'

export const Header = ({ startGame }) => (
    <header className="header">
        <img src={logo} alt="logo"/>
        <button className="btn" onClick={() => startGame()}> Start game </button>
    </header>
)