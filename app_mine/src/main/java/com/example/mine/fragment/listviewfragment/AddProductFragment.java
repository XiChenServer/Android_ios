package com.example.mine.fragment.listviewfragment;

import android.net.Uri;
import android.os.Bundle;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.fragment.app.FragmentManager;
import androidx.fragment.app.FragmentTransaction;
import androidx.recyclerview.widget.GridLayoutManager;

import com.example.common.room.GlideEngine;
import com.example.common.room.utils.UtilsWay;
import com.example.core_net_work.MyRetrofit;
import com.example.mine.R;
import com.example.mine.databinding.FragmentAddProductBinding;
import com.example.mine.fragment.listviewfragment.adapter.AddPicAdapter;
import com.example.mine.fragment.listviewfragment.entity.PicRecyclerViewEntity;
import com.example.mine.fragment.listviewfragment.eventbus.SendMessageEvent;
import com.github.jdsjlzx.interfaces.OnItemClickListener;
import com.github.jdsjlzx.recyclerview.LRecyclerViewAdapter;
import com.google.android.material.snackbar.Snackbar;
import com.luck.picture.lib.basic.PictureSelector;
import com.luck.picture.lib.config.SelectMimeType;
import com.luck.picture.lib.entity.LocalMedia;
import com.luck.picture.lib.interfaces.OnResultCallbackListener;
import com.tencent.mmkv.MMKV;

import org.greenrobot.eventbus.EventBus;
import org.greenrobot.eventbus.Subscribe;

import java.io.File;
import java.io.IOException;
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
 * @Date 2024-02-23 21:36
 * @Version 1.0
 */
public class AddProductFragment extends Fragment {


    @Override
    public void onStart() {
        super.onStart();
        EventBus.getDefault().register(this);
    }

    LRecyclerViewAdapter adapter;

    @Override
    public void onStop() {
        EventBus.getDefault().unregister(this);
        super.onStop();
    }

    FragmentAddProductBinding binding;
    List<PicRecyclerViewEntity> list = new ArrayList<>();

    @Subscribe
    public void isDelete(SendMessageEvent event) {
        if (event.isDelete()) {
            list.remove(event.getPosition());
            adapter.notifyDataSetChanged();
        }
    }

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_add_product, container, false);
        binding = FragmentAddProductBinding.bind(view);

        binding.transparentViewProduct.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                //不做任何操作，只是为了消化点击事件
            }
        });

        binding.chipIsAuction.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                if (binding.chipIsAuction.isChecked()) {
                    // 如果 Chip 被选中，则执行此处的逻辑
                    Toast.makeText(getActivity(), "Chip Checked", Toast.LENGTH_SHORT).show();
                } else {
                    //如果 Chip 被取消选中，则执行此处的逻辑
                    Toast.makeText(getActivity(), "Chip Unchecked", Toast.LENGTH_SHORT).show();
                }
            }
        });
        binding.chipClose.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                FragmentManager fragmentManager = getActivity().getSupportFragmentManager();
                FragmentTransaction fragmentTransaction = fragmentManager.beginTransaction();
                fragmentTransaction.setCustomAnimations(com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right, com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right);
                fragmentManager.popBackStack();
                fragmentTransaction.commit();
            }
        });

        binding.chipSend.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                binding.progressBarAddProduct.setVisibility(View.VISIBLE);
                binding.constraintAddProduct.setVisibility(View.GONE);
                MultipartBody.Builder builder = new MultipartBody.Builder().setType(MultipartBody.FORM);
                for (PicRecyclerViewEntity picRecyclerViewEntity : list) {
                    File file = picRecyclerViewEntity.getFile();
                    RequestBody body = RequestBody.create(MediaType.parse("multipart/form-data"), file);
                    builder.addFormDataPart("files", file.getName(), body); //添加图片数据，body创建的请求体
                }
                List<MultipartBody.Part> parts = builder.build().parts();
                List<RequestBody> myClass = new ArrayList<>();
                myClass.add(RequestBody.create(MediaType.parse("multipart/form-data"), "大象"));
                RequestBody requestFile_title = RequestBody.create(MediaType.parse("multipart/form-data"), binding.textInputEditText.getText().toString());
                RequestBody requestFile_number = RequestBody.create(MediaType.parse("multipart/form-data"), binding.editTextNumberPrice.getText().toString());
                RequestBody requestFile_information = RequestBody.create(MediaType.parse("multipart/form-data"), binding.textInputEditTextInfo.getText().toString());
                RequestBody requestFile_price = RequestBody.create(MediaType.parse("multipart/form-data"), binding.editTextNumberPrice.getText().toString());
                RequestBody requestFile_is_auction = null;
                try {
                    requestFile_is_auction = RequestBody.create(MediaType.parse("multipart/form-data"), binding.chipIsAuction.isChecked() ? "1" : "0");
                } catch (Exception e) {
                    throw new RuntimeException(e);
                }
                RequestBody requestFile_country = RequestBody.create(MediaType.parse("multipart/form-data"), "中国");
                RequestBody requestFile_province = RequestBody.create(MediaType.parse("multipart/form-data"), "陕西省");
                RequestBody requestFile_city = RequestBody.create(MediaType.parse("multipart/form-data"), "西安市");
                RequestBody requestFile_contact = RequestBody.create(MediaType.parse("multipart/form-data"), "");
                RequestBody requestFile_post_code = RequestBody.create(MediaType.parse("multipart/form-data"), "");

                MyRetrofit.serviceAPI.addProduct("bearer " + MMKV.defaultMMKV().getString("token", null), myClass, requestFile_title, requestFile_number, requestFile_information, requestFile_price, requestFile_is_auction, requestFile_country, requestFile_province, requestFile_city, requestFile_contact, requestFile_post_code, parts).enqueue(new Callback() {
                    @Override
                    public void onResponse(Call call, Response response) {
                        if (response.isSuccessful()) {
                            binding.progressBarAddProduct.setVisibility(View.GONE);
                            binding.constraintAddProduct.setVisibility(View.VISIBLE);
                            Snackbar.make(view, "上传成功", Snackbar.LENGTH_SHORT).show();
                            FragmentManager fragmentManager = getActivity().getSupportFragmentManager();
                            FragmentTransaction fragmentTransaction = fragmentManager.beginTransaction();
                            fragmentTransaction.setCustomAnimations(com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right, com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right);
                            fragmentManager.popBackStack();
                            fragmentTransaction.commit();
                        } else {
                            try {
                                Log.d("恶心的问题", response.errorBody().byteString().toString());
                            } catch (IOException e) {
                                Log.d("恶心的问题", e.toString());
                            }
                            Toast.makeText(getActivity(), "有问题", Toast.LENGTH_SHORT).show();
                        }
                    }

                    @Override
                    public void onFailure(Call call, Throwable t) {
                        Toast.makeText(getActivity(), "失败", Toast.LENGTH_SHORT).show();
                    }
                });
            }
        });
        adapter = new LRecyclerViewAdapter(new AddPicAdapter(list, getActivity()));
        binding.LRecyclerViewPickPic.setLayoutManager(new GridLayoutManager(getActivity(), 3));
        binding.LRecyclerViewPickPic.setAdapter(adapter);
        binding.LRecyclerViewPickPic.setPullRefreshEnabled(false);
        adapter.setOnItemClickListener(new OnItemClickListener() {
            @Override
            public void onItemClick(View view, int position) {
                if (list.size() == 0 || position == list.size()) {
                    PictureSelector.create(getActivity()).openGallery(SelectMimeType.ofImage()).setImageEngine(GlideEngine.createGlideEngine()).forResult(new OnResultCallbackListener<LocalMedia>() {
                        @Override
                        public void onResult(ArrayList<LocalMedia> result) {
                            for (LocalMedia localMedia : result) {
                                Log.d("文件的地址", localMedia.getAvailablePath());
                                list.add(new PicRecyclerViewEntity(new File(UtilsWay.getFilePathFromUri(getActivity(), Uri.parse(localMedia.getAvailablePath())))));
                            }
                            adapter.notifyDataSetChanged();
                        }

                        @Override
                        public void onCancel() {

                        }
                    });
                } else {
                    Bundle bundle = new Bundle();
                    bundle.putString("path", list.get(position).getFile().getAbsolutePath());
                    bundle.putInt("position", position);
                    FragmentManager fragmentManager = getActivity().getSupportFragmentManager();
                    FragmentTransaction fragmentTransaction = fragmentManager.beginTransaction();
                    ImagePreViewFragment fragment = new ImagePreViewFragment();
                    fragment.setArguments(bundle);
                    fragmentTransaction.setCustomAnimations(com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right, com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right);
                    fragmentTransaction.add(R.id.frame_mine, fragment);
                    fragmentTransaction.addToBackStack(null);
                    fragmentTransaction.commit();
                }
            }
        });

        return binding.getRoot();
    }
}
