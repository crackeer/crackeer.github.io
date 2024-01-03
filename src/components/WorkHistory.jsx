import React from 'react';
import { Collapse, Card } from 'antd';

const text = `
  A dog is a type of domesticated animal.
  Known for its loyalty and faithfulness,
  it can be found as a welcome guest in many households across the world.
`;

const items = [
  {
    key: '1',
    label: <strong >贝壳·如视 / Golang资深 / 2020.07 至 ~</strong>,
    children: <Card bodyStyle={{padding:10}}>{text}</Card>
  },
  {
    key: '2',
    label: 'This is panel header 2',
    children: <p>{text}</p>,
  },
  {
    key: '3',
    label: 'This is panel header 3',
    children: <p>{text}</p>,
  },
];

export default () => {
   return  <Collapse ghost defaultActiveKey={['1']} items={items} />
};