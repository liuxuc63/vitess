package com.flipkart.vitess.util;

import com.youtube.vitess.client.Context;
import com.youtube.vitess.proto.Topodata;
import com.youtube.vitess.proto.Vtrpc;
import org.joda.time.Duration;

/**
 * Created by naveen.nahata on 24/02/16.
 */
public class CommonUtils {
    /**
     * Create context used to create grpc client and executing query.
     *
     * @param username
     * @param connectionTimeout
     * @return
     */
    public static Context createContext(String username, long connectionTimeout) {
        Context context;
        Vtrpc.CallerID callerID = null;
        if (null != username) {
            callerID = Vtrpc.CallerID.newBuilder().setPrincipal(username).build();
        }
        if (null != callerID) {
            context = Context.getDefault().withDeadlineAfter(Duration.millis(connectionTimeout))
                .withCallerId(callerID);
        } else {
            context = Context.getDefault().withDeadlineAfter(Duration.millis(connectionTimeout));
        }
        return context;
    }

    /**
     * Converting string tabletType to Topodata.TabletType
     *
     * @param tabletType
     * @return
     */
    public static Topodata.TabletType getTabletType(String tabletType) {
        switch (tabletType.toLowerCase()) {
            case "master":
                return Topodata.TabletType.MASTER;
            case "replica":
                return Topodata.TabletType.REPLICA;
            case "rdonly":
                return Topodata.TabletType.RDONLY;
            default:
                return Topodata.TabletType.UNRECOGNIZED;
        }
    }

}

