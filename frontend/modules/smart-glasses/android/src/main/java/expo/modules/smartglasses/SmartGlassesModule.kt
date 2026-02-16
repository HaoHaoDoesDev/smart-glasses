package expo.modules.smartglasses

import expo.modules.kotlin.modules.Module
import expo.modules.kotlin.modules.ModuleDefinition
import com.qcwireless.bluetooth.BleScannerHelper
import com.qcwireless.bluetooth.BleOperateManager
import com.qcwireless.bluetooth.ScanWrapperCallback
import com.qcwireless.bluetooth.LargeDataHandler

class SmartGlassesModule : Module() {
  override fun definition() = ModuleDefinition {
    Name("SmartGlasses")

    Events("onDeviceFound", "onStatusUpdate", "onBatteryUpdate")

    Function("startScan") {
      val context = appContext.reactContext ?: return@Function
      
      BleScannerHelper.getInstance().scanDevice(context, null, object : ScanWrapperCallback() {
        override fun onScanResult(result: ScanResult) {
          sendEvent("onDeviceFound", mapOf(
            "name" to (result.device.name ?: "Unknown Glasses"),
            "address" to result.device.address,
            "rssi" to result.rssi
          ))
        }
      })
    }
    Function("stopScan") {
      val context = appContext.reactContext ?: return@Function
      BleScannerHelper.getInstance().stopScan(context)
    }

    Function("connect") { macAddress: String ->
      BleOperateManager.getInstance().connectDirectly(macAddress)
    }

    OnCreate {
      LargeDataHandler.getInstance().addOutDeviceListener(100, object : GlassesDeviceNotifyListener() {
        override fun parseData(cmdType: Int, response: GlassesDeviceNotifyRsp) {
          val command = response.loadData[6].toInt()
          
          // Relay battery updates back to the UI
          if (command == 0x05) { 
            sendEvent("onBatteryUpdate", mapOf(
              "level" to response.loadData[7].toInt(),
              "isCharging" to (response.loadData[8].toInt() == 1)
            ))
          }
        }
      })
    }
  }
}