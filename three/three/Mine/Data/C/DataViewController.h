//
//  DataViewController.h
//  three
//
//  Created by Huc on 2023/12/10.
//

#import <UIKit/UIKit.h>

NS_ASSUME_NONNULL_BEGIN

@interface DataViewController : UIViewController
<UITableViewDelegate, UITableViewDataSource>

@property UITableView *tableView;
@property NSArray *data;

@end

NS_ASSUME_NONNULL_END
