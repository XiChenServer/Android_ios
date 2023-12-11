//
//  AlterViewController.m
//  three
//
//  Created by Huc on 2023/12/11.
//

#import "AlterViewController.h"

@interface AlterViewController ()

@end

@implementation AlterViewController

- (void)viewDidLoad {
    [super viewDidLoad];
    [self loadBasicView];
    [self loadExtraView];
    // Do any additional setup after loading the view.
}

- (void)loadBasicView {
    self.view.backgroundColor = UIColor.systemGray6Color;
    
    self.mainTextField = [[UITextField alloc] init];
    self.subTextField = [[UITextField alloc] init];
    self.label = [[UILabel alloc] init];
    self.codeBtn = [UIButton buttonWithType:UIButtonTypeCustom];
    self.mainTextField.backgroundColor = UIColor.whiteColor;
    self.subTextField.backgroundColor = UIColor.whiteColor;
    self.codeBtn.backgroundColor = UIColor.systemBlueColor;
    [self.codeBtn setTintColor:UIColor.whiteColor];
    self.label.font = [UIFont fontWithName:@"Arial-BoldMT" size:30];
    
    self.mainTextField.frame = CGRectMake(self.view.bounds.size.width/2 - 150, self.view.bounds.size.height/2 - 100, 300, 60);
    self.subTextField.frame = CGRectMake(self.view.bounds.size.width/2 - 150, self.view.bounds.size.height/2, 300, 60);
    self.codeBtn.frame = CGRectMake(self.view.bounds.size.width/2 - 150, self.view.bounds.size.height/2 + 100, 300, 60);
    self.label.frame = CGRectMake(self.view.bounds.size.width/2 - 150, self.view.bounds.size.height/2 - 200, 300, 60);
}

- (void)loadExtraView {
    if (self.row == 1) {
        [self loadAlterNicknameView];
    } else if (self.row == 10) {
        [self loadAlterTelView];
    } else if (self.row == 11) {
        [self loadAlterEmailView];
    } else if (self.row == 12) {
        [self loadAlterWeChatView];
    } else if (self.row == 21) {
        [self loadAlterPasswordView];
    }
}

- (void)loadAlterNicknameView {
    self.mainTextField.text = @"Lai'";
    self.mainTextField.placeholder = @"输入昵称";
    [self.codeBtn setTitle:@"保存" forState:UIControlStateNormal];
    self.label.text = @"更改昵称";
    self.codeBtn.frame = self.subTextField.frame;
    [self.view addSubview:self.mainTextField];
    [self.view addSubview:self.codeBtn];
    [self.view addSubview:self.label];
}

- (void)loadAlterTelView {
    self.mainTextField.placeholder = @"输入手机号";
    self.subTextField.placeholder = @"输入验证码";
    [self.codeBtn setTitle:@"发送验证码" forState:UIControlStateNormal];
    self.label.text = @"绑定手机";
    [self.view addSubview:self.mainTextField];
    [self.view addSubview:self.subTextField];
    [self.view addSubview:self.codeBtn];
    [self.view addSubview:self.label];
}

- (void)loadAlterEmailView {
    self.mainTextField.placeholder = @"输入邮箱号";
    self.subTextField.placeholder = @"输入验证码";
    [self.codeBtn setTitle:@"发送验证码" forState:UIControlStateNormal];
    self.label.text = @"绑定邮箱";
    [self.view addSubview:self.mainTextField];
    [self.view addSubview:self.subTextField];
    [self.view addSubview:self.codeBtn];
    [self.view addSubview:self.label];
}

- (void)loadAlterWeChatView {
    self.mainTextField.placeholder = @"输入微信号";
    self.subTextField.placeholder = @"输入验证码";
    [self.codeBtn setTitle:@"发送验证码" forState:UIControlStateNormal];
    self.label.text = @"绑定微信";
    [self.view addSubview:self.mainTextField];
    [self.view addSubview:self.subTextField];
    [self.view addSubview:self.codeBtn];
    [self.view addSubview:self.label];
}

- (void)loadAlterPasswordView {
    self.mainTextField.placeholder = @"输入新密码";
    self.subTextField.placeholder = @"重复新密码";
    self.mainTextField.secureTextEntry = YES;
    self.subTextField.secureTextEntry = YES;
    [self.codeBtn setTitle:@"保存" forState:UIControlStateNormal];
    self.label.text = @"修改密码";
    [self.view addSubview:self.mainTextField];
    [self.view addSubview:self.subTextField];
    [self.view addSubview:self.codeBtn];
    [self.view addSubview:self.label];
}

- (void)alterEmail {
    
}

- (void)send {
    NSURL *url = [NSURL URLWithString:@"http://192.168.0.146:9000/send_phone_code"];
    NSMutableURLRequest *request = [NSMutableURLRequest requestWithURL:url];
    request.HTTPMethod = @"POST";

    // 创建要发送的JSON参数
    NSDictionary *parameters = @{@"phone": self.mainTextField.text};
    NSError *error;
    NSData *jsonData = [NSJSONSerialization dataWithJSONObject:parameters options:0 error:&error];
    if (error) {
        NSLog(@"Failed to serialize JSON: %@", error);
        return;
    }

    // 设置请求体为JSON参数
    request.HTTPBody = jsonData;
    [request setValue:@"application/json" forHTTPHeaderField:@"Content-Type"];

    NSURLSession *session = [NSURLSession sharedSession];
    NSURLSessionDataTask *dataTask = [session dataTaskWithRequest:request completionHandler:^(NSData *data, NSURLResponse *response, NSError *error) {
        if (error) {
            NSLog(@"Failed to send request: %@", error);
            return;
        }
        
        // 处理响应数据
        NSString *responseData = [[NSString alloc] initWithData:data encoding:NSUTF8StringEncoding];
        NSLog(@"Response: %@", responseData);
    }];

    [dataTask resume];
}

/*
#pragma mark - Navigation

// In a storyboard-based application, you will often want to do a little preparation before navigation
- (void)prepareForSegue:(UIStoryboardSegue *)segue sender:(id)sender {
    // Get the new view controller using [segue destinationViewController].
    // Pass the selected object to the new view controller.
}
*/

@end
