import React from 'react'

export const Letter = ({letter, width, enter, click, selection}) => (
    <div 
        className={`letter ${selection ? 'selected' : ''}`} 
        style={{width: width+'%'}}
    >
        <p onMouseEnter={() => enter()} onMouseDown={() => click()}>{ letter }</p>
    </div>
)
