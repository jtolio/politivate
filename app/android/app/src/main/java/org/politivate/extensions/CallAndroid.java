package org.politivate.extensions;

import android.app.Activity;
import android.content.Intent;
import android.net.Uri;

import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;

public class CallAndroid extends ReactContextBaseJavaModule {

  public CallAndroid(ReactApplicationContext reactContext) {
    super(reactContext);
  }

  @Override
  public String getName() {
    return "CallAndroid";
  }

  @ReactMethod
  public void call(String phonenumber) {
    Activity activity = getCurrentActivity();
    if (activity == null) {
      return;
    }
    Intent intent = new Intent(Intent.ACTION_CALL);
    intent.setData(Uri.parse("tel:" + phonenumber));
    activity.startActivity(intent);
  }
}
