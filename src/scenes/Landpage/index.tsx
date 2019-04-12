import React, { Component } from 'react';
import './style.scss';
import { Navigation } from '../../components/Navigation';

export class Landpage extends Component {
  public render() {
    const links = [

      {
        header: 'Facebook',
        link: 'https://www.facebook.com/ieeeIonianUni/',
        newTab: true
      },
      {
        header: 'Instagram',
        link: 'https://www.instagram.com/iusbieee/',
        newTab: true
      }
    ];
    return (
      <div className='landpage'>
        <div className='container'>
          <header>
            <h1>
              Ionian IEEE Student Branch
            </h1>
          </header>
          <Navigation links={links} reactRouter={false} horizontal/>
        </div>
      </div>
    );
  }
}
