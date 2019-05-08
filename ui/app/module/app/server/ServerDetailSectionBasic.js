import React from "react";
import PropTypes from "prop-types";
import CSSModules from "react-css-modules";
import styles from "./App.css";

function ServerDetailSectionBasic(props) {
  return (
    <div styleName="ServerDetailSectionDiv">
      <table>
        <thead>
          <tr>
            <th colSpan="2" styleName="level1">
              标识
            </th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <th styleName="level2">ID</th>
            <td>{props.server.ID}</td>
          </tr>
          <tr>
            <th styleName="level2">名称</th>
            <td>{props.server.Name}</td>
          </tr>
          <tr>
            <th styleName="level2">描述</th>
            <td>{props.server.Description}</td>
          </tr>
          <tr>
            <th styleName="level2">类型</th>
            <td>{props.server.Type}</td>
          </tr>
        </tbody>
        <thead>
          <tr>
            <th colSpan="2" styleName="level1">
              状态
            </th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <th styleName="level2">状态</th>
            <td>{props.server.State}</td>
          </tr>
          <tr>
            <th styleName="level2">健康</th>
            <td>{props.server.Health}</td>
          </tr>
          <tr>
            <th styleName="level2">电源</th>
            <td>{props.server.PowerState}</td>
          </tr>
          <tr>
            <th styleName="level2">LED</th>
            <td>{props.server.IndicatorLED}</td>
          </tr>
        </tbody>
        <thead>
          <tr>
            <th colSpan="2" styleName="level1">
              资产信息
            </th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <th styleName="level2">制造商</th>
            <td>{props.server.Manufacturer}</td>
          </tr>
          <tr>
            <th styleName="level2">类型</th>
            <td>{props.server.Model}</td>
          </tr>
          <tr>
            <th styleName="level2">库存量单位</th>
            <td>{props.server.SKU}</td>
          </tr>
          <tr>
            <th styleName="level2">序列号</th>
            <td>{props.server.SerialNumber}</td>
          </tr>
          <tr>
            <th styleName="level2">零件号</th>
            <td>{props.server.PartNumber}</td>
          </tr>
          <tr>
            <th styleName="level2">备件号</th>
            <td>{props.server.SparePartNumber}</td>
          </tr>
          <tr>
            <th styleName="level2">资产标签</th>
            <td>{props.server.AssetTag}</td>
          </tr>
        </tbody>
      </table>
    </div>
  );
}

ServerDetailSectionBasic.propTypes = {
  server: PropTypes.object
};

export default CSSModules(ServerDetailSectionBasic, styles);
