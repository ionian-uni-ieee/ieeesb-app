import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import './style.scss';

interface Props {
  links: LinkProp[];
  reactRouter: boolean;
  horizontal: boolean;
}

interface LinkProp {
  header: string | JSX.Element;
  link: string;
  newTab: boolean;
}

interface ListStyle {
  width: string;
  height: string;
  background: string;
  flexDirection: 'column' | 'row';
  flexFlow: 'column' | 'row';
}

interface ListElementStyle {
  width: string;
  height: string;
}

export class Navigation extends Component<Props> {
  listStyle: ListStyle = {
    width: 'auto',
    height: 'auto',
    background: 'linear-gradient(120deg, #dd1155, #bb1122)',
    flexDirection: 'column',
    flexFlow: 'column'
  };
  listElementStyle: ListElementStyle = {
    height: '100%',
    width: 'auto'
  }

  constructor(props: Props) {
    super(props);

    if (this.props.horizontal) {
      this.listStyle.width = '100%';
      this.listStyle.height = '4em';
      this.listStyle.flexDirection = 'row';
      this.listStyle.flexFlow = 'row';
      // this.listElementStyle.width = `${100 / this.props.links.length}%`;
      // this.listElementStyle.height = '100%';
    } else {
      this.listStyle.height = '100%'; 
      this.listStyle.width = '4em';
      // this.listElementStyle.height = `${100 / this.props.links.length}%`;
      // this.listElementStyle.width = '100%';
    }
  }

  getRenderedLinkComponent(header: string | JSX.Element, link: string, newTab: boolean): JSX.Element {
    const component = this.props.reactRouter ? <Link to={link} >{header}</Link>: <a href={link} target={newTab ? '_blank' : ''}>{header}</a>;
    return (
      component
    );
  }

  getRenderedLinks(): JSX.Element {
    const links = this.props.links.map((link) => {
      return (
        <li id={link.header.toString()} style={this.listElementStyle}>
          {this.getRenderedLinkComponent(link.header, link.link, link.newTab)}
        </li>
      );
    });
    return (
      <ul style={this.listStyle}>
        {links}
      </ul>
    );
  }

  render() {
    return (
      <div className='navigation' >
        {this.getRenderedLinks()}
      </div>
    );
  }
}
