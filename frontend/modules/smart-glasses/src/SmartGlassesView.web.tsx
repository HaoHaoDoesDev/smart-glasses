import * as React from 'react';

import { SmartGlassesViewProps } from './SmartGlasses.types';

export default function SmartGlassesView(props: SmartGlassesViewProps) {
  return (
    <div>
      <iframe
        style={{ flex: 1 }}
        src={props.url}
        onLoad={() => props.onLoad({ nativeEvent: { url: props.url } })}
      />
    </div>
  );
}
