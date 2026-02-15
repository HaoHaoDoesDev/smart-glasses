import { requireNativeView } from 'expo';
import * as React from 'react';

import { SmartGlassesViewProps } from './SmartGlasses.types';

const NativeView: React.ComponentType<SmartGlassesViewProps> =
  requireNativeView('SmartGlasses');

export default function SmartGlassesView(props: SmartGlassesViewProps) {
  return <NativeView {...props} />;
}
