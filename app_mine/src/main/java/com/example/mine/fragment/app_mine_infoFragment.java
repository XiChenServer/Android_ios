package com.example.mine.fragment;

import android.content.Context;
import android.graphics.Bitmap;
import android.graphics.drawable.Drawable;
import android.net.Uri;
import android.os.Bundle;
import android.os.Handler;
import android.os.HandlerThread;
import android.text.InputType;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageView;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.fragment.app.FragmentManager;
import androidx.fragment.app.FragmentTransaction;

import com.afollestad.materialdialogs.MaterialDialog;
import com.alibaba.android.arouter.facade.annotation.Route;
import com.bumptech.glide.Glide;
import com.bumptech.glide.request.target.CustomTarget;
import com.bumptech.glide.request.transition.Transition;
import com.example.common.room.AppDatabase;
import com.example.common.room.GlideEngine;
import com.example.common.room.entitues.User;
import com.example.core_net_work.MyRetrofit;
import com.example.core_net_work.model.userInfo.ChangeUserInfoRequest;
import com.example.core_net_work.model.userInfo.ChangeUserInfoResult;
import com.example.core_net_work.model.userInfo.UploadAvatarResult;
import com.example.mine.R;
import com.example.mine.databinding.FragmentAppMineInfoBinding;
import com.google.android.material.snackbar.Snackbar;
import com.luck.picture.lib.basic.PictureSelector;
import com.luck.picture.lib.config.SelectMimeType;
import com.luck.picture.lib.engine.CropFileEngine;
import com.luck.picture.lib.entity.LocalMedia;
import com.luck.picture.lib.interfaces.OnResultCallbackListener;
import com.tencent.mmkv.MMKV;
import com.yalantis.ucrop.UCrop;
import com.yalantis.ucrop.UCropImageEngine;

import java.io.File;
import java.util.ArrayList;
import java.util.List;

import okhttp3.MediaType;
import okhttp3.MultipartBody;
import okhttp3.RequestBody;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

/**
 * @Author winiymissl
 * @Date 2023-12-17 20:55
 * @Version 1.0
 */
@Route(path = "/mine/fragment/app_mine_infoFragment")
public class app_mine_infoFragment extends Fragment {
    FragmentAppMineInfoBinding binding;


    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_app_mine_info, container, false);
        binding = FragmentAppMineInfoBinding.bind(view);
        //从数据库中加载数据
        AppDatabase instance = AppDatabase.getInstance(getActivity());
        if (instance.isOpen()) {
            HandlerThread handlerThread = new HandlerThread("DatabaseThread");
            handlerThread.start();
            Handler handler = new Handler(handlerThread.getLooper());
            handler.post(new Runnable() {
                @Override
                public void run() {
                    // 在这里执行数据库操作
                    // 注意：不要在这里更新 UI，因为这是在后台线程上执行的操作
                    List<User> allInfo = instance.userDao().getAllInfo();
                    User user = allInfo.get(0);
                    //主线程更新ui
                    getActivity().runOnUiThread(new Runnable() {
                        @Override
                        public void run() {
                            binding.textViewEmailEmail.setText(user.getEmail());
                            binding.textViewWechatWechat.setText(user.getWechat_number());
                            binding.textViewNameName.setText(user.getNickname());
                            binding.textViewPhone.setText(user.phone_number);
                            binding.textViewAccountAccount.setText(user.getAccount());
                            if (user.getAddress() != null) {
                                binding.textViewAddressAddress.setText("" + user.getAddress().getCountry() + " " + user.getAddress().getProvince() + " " + user.getAddress().getCity() + ".");
                            } else {
                                binding.textViewAddressAddress.setText("");
                            }
                            Glide.with(getActivity()).load(user.getAvatar()).error(com.example.common.R.drawable.avatatloadfail).into(binding.imageViewAvator);
                        }
                    });
                }
            });
        }

        binding.constraintAvator.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                PictureSelector.create(getActivity()).openGallery(SelectMimeType.ofImage()).setCropEngine(new CropFileEngine() {
                    @Override
                    public void onStartCrop(Fragment fragment, Uri srcUri, Uri destinationUri, ArrayList<String> dataSource, int requestCode) {
                        UCrop uCrop = UCrop.of(srcUri, destinationUri, dataSource);
                        UCrop.Options options = new UCrop.Options();
                        options.isDarkStatusBarBlack(true);
                        options.isForbidSkipMultipleCrop(true);
                        uCrop = uCrop.withAspectRatio(1, 1).withOptions(options);
                        uCrop.setImageEngine(new UCropImageEngine() {
                            @Override
                            public void loadImage(Context context, String url, ImageView imageView) {
                                Glide.with(context).load(url).into(imageView);
                            }

                            @Override
                            public void loadImage(Context context, Uri url, int maxWidth, int maxHeight, OnCallbackListener<Bitmap> call) {
                                Glide.with(context).asBitmap().override(maxWidth, maxHeight).load(url).into(new CustomTarget<Bitmap>() {
                                    @Override
                                    public void onResourceReady(@NonNull Bitmap resource, @Nullable Transition<? super Bitmap> transition) {
                                        if (call != null) {
                                            call.onCall(resource);
                                        }
                                    }

                                    @Override
                                    public void onLoadFailed(@Nullable Drawable errorDrawable) {
                                        if (call != null) {
                                            call.onCall(null);
                                        }
                                    }

                                    @Override
                                    public void onLoadCleared(@Nullable Drawable placeholder) {
                                    }
                                });
                            }
                        });
                        uCrop.start(fragment.getActivity(), fragment, requestCode);
                    }
                }).setMaxSelectNum(1).setImageEngine(GlideEngine.createGlideEngine()).forResult(new OnResultCallbackListener<LocalMedia>() {
                    @Override
                    public void onResult(ArrayList<LocalMedia> result) {
                        Log.d("问题", result.get(0).getAvailablePath());
                        File file = new File(result.get(0).getCutPath());
                        RequestBody fileRQ = RequestBody.create(MediaType.parse("image/*"), file);

                        MyRetrofit.serviceAPI.uploadAvatar("Bearer " + MMKV.defaultMMKV().getString("token", null), MultipartBody.Part.createFormData("files", file.getName(), fileRQ)).enqueue(new Callback<UploadAvatarResult>() {

                            @Override
                            public void onResponse(Call<UploadAvatarResult> call, Response<UploadAvatarResult> response) {
                                if (response.isSuccessful()) {
                                    Snackbar snackbar = Snackbar.make(view, "上传成功", Snackbar.LENGTH_SHORT);
                                    snackbar.show();
                                } else {
                                    Snackbar snackbar = Snackbar.make(view, "上传有问题", Snackbar.LENGTH_SHORT);
                                    snackbar.show();
                                }
                                Glide.with(getActivity()).load(file).placeholder(com.example.common.R.drawable.loading).error(com.example.common.R.drawable.avatatloadfail).into(binding.imageViewAvator);
                            }

                            @Override
                            public void onFailure(Call<UploadAvatarResult> call, Throwable t) {
                                Snackbar snackbar = Snackbar.make(requireView(), "上传失败", Snackbar.LENGTH_SHORT);
                                snackbar.show();
                            }
                        });
                    }

                    @Override
                    public void onCancel() {

                    }
                });
            }
        });
        binding.transparentView.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                //防止随意点击
            }
        });
        binding.constraintPhone.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                new MaterialDialog.Builder(getActivity()).title("电话").inputType(InputType.TYPE_CLASS_TEXT).input("新电话", "", new MaterialDialog.InputCallback() {
                    @Override
                    public void onInput(@NonNull MaterialDialog dialog, CharSequence input) {
                        Bundle bundle = new Bundle();
                        bundle.putString("phone", input.toString());
                        FragmentManager fragmentManager = getActivity().getSupportFragmentManager();
                        FragmentTransaction fragmentTransaction = fragmentManager.beginTransaction();
                        app_mine_savePhoneFragment fragment = new app_mine_savePhoneFragment();
                        fragment.setArguments(bundle);
                        fragmentTransaction.setCustomAnimations(com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right, com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right);
                        fragmentTransaction.add(R.id.frame_mine, fragment);
                        fragmentTransaction.addToBackStack(null);
                        fragmentTransaction.commit();
                    }
                }).show();
            }
        });
        Log.d("问题", MMKV.defaultMMKV().getString("token", null));
        binding.cosntraintName.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                //设置名字，使用inputDialog
                new MaterialDialog.Builder(getActivity()).title("名字").inputType(InputType.TYPE_CLASS_TEXT).input("name", binding.textViewNameName.getText().toString(), new MaterialDialog.InputCallback() {
                    @Override
                    public void onInput(MaterialDialog dialog, CharSequence input) {
                        if (!input.toString().equals("")) {
                            HandlerThread handlerThread = new HandlerThread("DatabaseThread");
                            handlerThread.start();
                            Handler handler = new Handler(handlerThread.getLooper());
                            handler.post(new Runnable() {

                                @Override
                                public void run() {
                                    List<User> allInfo = AppDatabase.getInstance(getActivity()).userDao().getAllInfo();
                                    User user = allInfo.get(0);
                                    MyRetrofit.serviceAPI.changeUserInfo("bearer " + MMKV.defaultMMKV().getString("token", null), new ChangeUserInfoRequest(user.getName(), input.toString(), user.getWechat_number(), user.getEmail())).enqueue(new Callback<ChangeUserInfoResult>() {
                                        @Override
                                        public void onResponse(Call<ChangeUserInfoResult> call, Response<ChangeUserInfoResult> response) {
                                            if (response.isSuccessful()) {
                                                Snackbar.make(view, "修改成功", Snackbar.LENGTH_SHORT).show();
                                                HandlerThread myThread = new HandlerThread("updateName");
                                                myThread.start();
                                                Handler myHandler = new Handler(myThread.getLooper());
                                                myHandler.post(new Runnable() {
                                                    @Override
                                                    public void run() {
                                                        AppDatabase.getInstance(getActivity()).userDao().changeUserInfo(binding.textViewNameName.getText().toString());
                                                    }
                                                });
                                            } else {
                                                Snackbar.make(view, "修改有问题", Snackbar.LENGTH_SHORT).show();
                                            }
                                        }

                                        @Override
                                        public void onFailure(Call<ChangeUserInfoResult> call, Throwable t) {
                                            Snackbar.make(view, "修改失败", Snackbar.LENGTH_SHORT).show();
                                        }
                                    });

                                }
                            });
                        }
                        binding.textViewNameName.setText(input);
                    }
                }).show();

            }
        });

//        binding.buttonSave.setOnClickListener(new View.OnClickListener() {
//            @Override
//            public void onClick(View v) {
//                FragmentManager fragmentManager = getActivity().getSupportFragmentManager();
//                FragmentTransaction fragmentTransaction = fragmentManager.beginTransaction();
//                app_mine_saveInfoFragment fragment = new app_mine_saveInfoFragment();
//                Bundle bundle = new Bundle();
//                bundle.putString("nickname", binding.textViewNameName.getText().toString());
//                bundle.putString("phone", binding.textViewPhone.getText().toString());
//                fragment.setArguments(bundle);
//                fragmentTransaction.setCustomAnimations(com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right, com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right);
//                fragmentTransaction.add(R.id.frame_mine, fragment);
//                fragmentTransaction.addToBackStack(null);
//                fragmentTransaction.commit();
//            }
//        });

        return view;
    }
}
