// Reexport the native module. On web, it will be resolved to SmartGlassesModule.web.ts
// and on native platforms to SmartGlassesModule.ts
export { default } from './src/SmartGlassesModule';
export { default as SmartGlassesView } from './src/SmartGlassesView';
export * from  './src/SmartGlasses.types';
