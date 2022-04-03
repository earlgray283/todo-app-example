import React, { useState } from 'react';

export const InputIdButton = (props: {
  label: string;
  onClick: (id: number) => void;
}): JSX.Element => {
  const [id, setID] = useState(0);
  return (
    <div>
      <input
        name='id'
        type='number'
        value={id}
        onChange={(e) => setID(parseInt(e.target.value, 10))}
      />
      <button onClick={() => props.onClick(id)}>{props.label}</button>
    </div>
  );
};
