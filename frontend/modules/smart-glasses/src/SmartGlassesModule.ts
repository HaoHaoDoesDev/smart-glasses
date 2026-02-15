import { NativeModule, requireNativeModule } from 'expo';

import { SmartGlassesModuleEvents } from './SmartGlasses.types';

declare class SmartGlassesModule extends NativeModule<SmartGlassesModuleEvents> {
  PI: number;
  hello(): string;
  setValueAsync(value: string): Promise<void>;
}

// This call loads the native module object from the JSI.
export default requireNativeModule<SmartGlassesModule>('SmartGlasses');
