import { registerWebModule, NativeModule } from 'expo';

import { ChangeEventPayload } from './SmartGlasses.types';

type SmartGlassesModuleEvents = {
  onChange: (params: ChangeEventPayload) => void;
}

class SmartGlassesModule extends NativeModule<SmartGlassesModuleEvents> {
  PI = Math.PI;
  async setValueAsync(value: string): Promise<void> {
    this.emit('onChange', { value });
  }
  hello() {
    return 'Hello world! 👋';
  }
};

export default registerWebModule(SmartGlassesModule, 'SmartGlassesModule');
