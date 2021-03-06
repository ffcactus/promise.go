import React from 'react';
import CSSModules from 'react-css-modules';
import styles from './App.css';

class IDPoolList extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div styleName="flex-item flex-row-container left-border flex-item-last">
        <p>ID Pool</p>
      </div>
    );
  }
}

export default CSSModules(IDPoolList, styles, {allowMultiple: true});
