import React from 'react';
import Description from './Description';

export default ({
                  name,
                  param,
                  opRef,
                  value,
                }) => {
  return (
    <div className='form-group'>
      <label className='form-control-label' htmlFor={name}>{name}</label>
      <Description value={param.description} opRef={opRef}/>
      <input
        className='form-control'
        id={name}
        readOnly={true}
        type={param.isSecret ? 'password' : 'number'}
        value={value || param.default || ""}
      />
    </div>
  );
}