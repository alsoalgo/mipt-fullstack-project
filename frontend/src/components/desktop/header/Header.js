import './Header.css';
import React from 'react';
import Logo from './logo/Logo';
import Menu from './menu/Menu';
import { Link } from 'react-router-dom';

export function Header() {
    return (
        <div className='header'>
            <Link to="/" className='header-logo'><Logo /></Link>
            <Menu />
        </div>
    );
}
